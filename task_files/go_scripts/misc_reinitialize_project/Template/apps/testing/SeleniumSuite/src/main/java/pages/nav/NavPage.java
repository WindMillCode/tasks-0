package pages.nav;

import baseclasses.BasePage;
import org.openqa.selenium.By;

public class NavPage extends BasePage {

  public static final String PARENTCLASS = "Nav";
  private NavActController act;
  private NavVerifyController verify;

  public NavActController act() {
    return act;
  }

  public NavVerifyController verify() {
    return verify;
  }

  private NavPage() {
    // hide it
  }

  private NavPage(NavActController act, NavVerifyController verify) {
    this.act = act;
    this.verify = verify;
  }

  public static NavPage getNavPage() {
    return new NavPage(new NavActController(), new NavVerifyController());
  }

  public static By pricingOption() {
    return By.id(generateCSSSelector(PARENTCLASS, "pricingOption"));
  }

  public static By storeOption() {
    return By.id(generateCSSSelector(PARENTCLASS, "storeOption"));
  }


  public static By plansOption() {
    return By.id(generateCSSSelector(PARENTCLASS, "plansOption"));
  }

  public static By productsOption() {
    return By.id(generateCSSSelector(PARENTCLASS, "productsOption"));
  }



}
