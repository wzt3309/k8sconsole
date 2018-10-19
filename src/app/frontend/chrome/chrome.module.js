import chromeDirective from "./chrome.directive";

export default angular.module(
    'k8sconsole.chrome',
    [
        'ngMaterial',
    ])
    .directive('chrome', chromeDirective);