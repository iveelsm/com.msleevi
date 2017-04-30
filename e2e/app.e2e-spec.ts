import { MsleeviPage } from './app.po';

describe('msleevi App', () => {
  let page: MsleeviPage;

  beforeEach(() => {
    page = new MsleeviPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
