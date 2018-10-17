export class MainPage {
    constructor() {
        this.jumbEl = element(by.css('.jumbotron'));
        this.h1El = this.jumbEl.element(by.css('h1'));
    }
}