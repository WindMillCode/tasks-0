package e2e;

import org.testng.annotations.Test;
import baseclasses.UIBaseClass;
import pages.nav.NavPage;
import util.E2EUtils;

public class TemplateTest extends UIBaseClass {
    NavPage nav = NavPage.getNavPage();
    E2EUtils e2eutil = E2EUtils.getCommonUtils();

    @Test
    public void navToProducts() throws InterruptedException {


    driver.get("https://windmillcode.com");

    e2eutil
      .waitForScreenToUpdate()
      .moveMouseToElement(driver, NavPage.pricingOption())
      .moveMouseToElement(driver, NavPage.storeOption())
      .click(driver, NavPage.productsOption())
      .waitForScreenToUpdate();

    String origPageTitlePrefix = E2EUtils.PAGE_TITLE_PREFIX;
    E2EUtils.PAGE_TITLE_PREFIX ="Windmillcode";
    e2eUtil
      .verifyPageTitle(driver, "Products");
    E2EUtils.PAGE_TITLE_PREFIX = origPageTitlePrefix;
  }

}
