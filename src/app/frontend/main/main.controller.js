export class MainController {
    constructor() {
        /** @export */
        this.testValue = 7;
    }

    /**
     * @param {!angular.$timeout} $timeout
     * @ngInject
     * @export
     */
    activate($timeout) {
        $timeout(() => {
            this.testValue = 8;
        }, 4000);
    }
}
