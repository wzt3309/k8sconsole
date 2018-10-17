import {MainController} from 'main/main.controller';

describe('Main controller', () => {
    let vm;

    beforeEach(() => {
        vm = new MainController();
    });

    it('should do something', () => {
        expect(vm.testValue).toEqual(7);
    });
});
