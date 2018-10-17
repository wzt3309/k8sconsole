export function routerConfig($componentLoaderProvider) {
    $componentLoaderProvider.setTemplateMapping(function(name) {
        return `${ name }/${ name }.html`;
    });
}

export class RouterController {
    /** @ngInject */
    constructor($router) {
        var router = $router;
        router['config']([
            { path: '/', component: 'main' }
        ]);
    }
}
