import {MainController} from "./main.controller";

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
    $stateProvider.state('main', {
        url: '',
        templateUrl: 'main/main.html',
        controller: MainController,
        controllerAs: 'ctrl',
    });
}