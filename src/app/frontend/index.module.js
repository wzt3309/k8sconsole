import mainModule from './main/main.module';
import routerConfig from './index.route';
import chromeModule from './chrome/chrome.module';

export default angular.module(
    'k8sconsole',
    [
        'ngAnimate',
        'ngAria',
        'ngMaterial',
        'ngMessages',
        'ngResource',
        'ngSanitize',
        'ui.router',
        mainModule.name,
        chromeModule.name,
    ])
    // .config(config)
    .config(routerConfig)
    // .run(runBlock)
    // .controller('RouterController', RouterController)
    // .controller('MainController', MainController);