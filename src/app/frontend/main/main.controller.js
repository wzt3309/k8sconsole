export class MainController {
    constructor() {
        this.testValue = 7;
    }

    activate($timeout) {
        $timeout(() => {
            this.foo = 'bar';
        }, 4000);
    }
}