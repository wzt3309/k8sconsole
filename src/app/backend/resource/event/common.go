package event

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/resource/common"
	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	client "k8s.io/client-go/kubernetes"
)

var EmptyEventList = &common.EventList{
	ListMeta: api.ListMeta{
		TotalItems: 0,
	},
	Events: make([]common.Event, 0),
}

// GetEvents gets events associated to resource with given name.
func GetEvents(client client.Interface, namespace, resourceName string) ([]v1.Event, error) {
	fieldSelector, err := fields.ParseSelector("involvedObject.name" + "=" + resourceName)
	if err != nil {
		return nil, err
	}

	channels := &common.ResourceChannels{
		EventList: common.GetEventListChannelWithOptions(
			client,
			common.NewOneNamespaceQuery(namespace),
			metaV1.ListOptions{
				LabelSelector: labels.Everything().String(),
				FieldSelector: fieldSelector.String(),
			},
			1),
	}

	eventList := <- channels.EventList.List
	if err := <- channels.EventList.Error; err != nil {
		return nil, err
	}

	return FillEventsType(eventList.Items), nil
}

// GetPodsEvents gets events targeting given list of pods.
func GetPodsEvents(client client.Interface, namespace string, pods []v1.Pod) ([]v1.Event, error) {
	nsQuery := common.NewOneNamespaceQuery(namespace)
	if namespace == v1.NamespaceAll {
		nsQuery = common.NewNamespaceQuery([]string{})
	}

	channels := &common.ResourceChannels{
		EventList: common.GetEventListChannel(client, nsQuery, 1),
	}

	eventList := <- channels.EventList.List
	if err := <- channels.EventList.Error; err != nil {
		return nil, err
	}

	events := filterEventsByPodsUID(eventList.Items, pods)
	return events, nil
}

// GetPodEvents gets pods events associated to pod name and namespace
func GetPodEvents(client client.Interface, namespace, podName string) ([]v1.Event, error) {
	channels := &common.ResourceChannels{
		PodList: common.GetPodListChannel(client,
			common.NewOneNamespaceQuery(namespace),
			1),
		EventList: common.GetEventListChannel(client, common.NewOneNamespaceQuery(namespace), 1),
	}

	podList := <- channels.PodList.List
	if err := <- channels.PodList.Error; err != nil {
		return nil, err
	}

	eventList := <- channels.EventList.List
	if err := <- channels.EventList.Error; err != nil {
		return nil, err
	}

	l := make([]v1.Pod, 0)
	for _, pod := range podList.Items {
		if pod.Name == podName {
			l = append(l, pod)
		}
	}

	events := filterEventsByPodsUID(eventList.Items, l)
	return FillEventsType(events), nil
}

func FillEventsType(events []v1.Event) []v1.Event {
	for i := range events {
		// fill in if the event with empty type
		if len(events[i].Type) == 0 {
			if isFailedReason(events[i].Reason, FailedReasonPartials...) {
				events[i].Type = v1.EventTypeWarning
			} else {
				events[i].Type = v1.EventTypeNormal
			}
		}
	}

	return events
}