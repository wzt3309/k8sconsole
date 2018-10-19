import ChromeController from "./chrome.controller";

/**
 * Returns directive definition object for the chrome directive.
 *
 * @return {!angular.Directive}
 */
export default function chromeDirective() {
    return {
        bindToController: true,
        controller: ChromeController,
        controllerAs: 'ctrl',
        templateUrl: 'chrome/chrome.html',
        transclude: true,
    };
}