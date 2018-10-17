import { config } from './index.config';
import { RouterController, routerConfig } from './index.route';
import { runBlock } from './index.run';
import { MainController } from './main/main.controller';

export default angular.module(
    'kubernetesConsole',
    ['ngAnimate', 'ngSanitize', 'ngMessages', 'ngAria', 'ngResource', 'ngNewRouter', 'ngMaterial'])
    .config(config)
    .config(routerConfig)
    .run(runBlock)
    .controller('RouterController', RouterController)
    .controller('MainController', MainController);