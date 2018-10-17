/**
 * @param {!ngNewRouter.$componentLoaderProvider} $componentLoaderProvider
 * @ngInject
 */
export function routerConfig($componentLoaderProvider) {
    $componentLoaderProvider.setTemplateMapping(function(name) {
        return `${name}/${name}.html`;
    });
}


export class RouterController {
    /**
     * @param {!ngNewRouter.$router} $router
     * @ngInject
     */
    constructor($router) {
        $router.config([
            { path: '/', component: 'main' }
        ]);
    }
}
