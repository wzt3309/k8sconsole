import {MainPage} from './main.po';


describe('The main view', function () {
    var page;

    beforeEach(function () {
        browser.get('/index.html');
        page = new MainPage();
    });

    it('should do something', function() {
        expect(page.h1El.getText()).toBe('\'Allo, \'Allo!');
    });
});