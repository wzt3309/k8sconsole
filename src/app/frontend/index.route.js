/**
 * Global route configuration for the application.
 *
 * @param {!ui.router.$urlRouterProvider} $urlRouterProvider
 * @ngInject
 */
export default function routerConfig($urlRouterProvider) {
    $urlRouterProvider.otherwise('');
}