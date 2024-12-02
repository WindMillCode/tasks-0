package automate;

import java.util.ArrayList;
import java.util.List;
import org.testng.annotations.Test;

import baseclasses.UIBaseClass;


public class MyAutomate extends UIBaseClass {


  // Example
  @Test
  public void workWithCheckOutPage() throws InterruptedException {
    ArrayList<Integer> colorOptions = new ArrayList<>(List.of(0, 1));
    ArrayList<Integer> sizeOptions = new ArrayList<>(List.of(0, 1, 2, 3, 4));

    // e2eUtil
    //   .waitForScreenToUpdate()
    //   .moveMouseToElement(driver, NavPage.pricingOption())
    //   .moveMouseToElement(driver, NavPage.storeOption())
    //   .click(driver, NavPage.productsOption())
    //   .clickFromOptions(driver, ProductsPage.selectProductCard(), 0)
    //   .enterIntoInput(driver, ProductdetailPage.quantityField(), "5")
    //   .clickFromOptions(driver, ProductdetailPage.colorFields(), CommonUtils.chooseRandomOptionFromSequence(colorOptions))
    //   .clickFromOptions(driver, ProductdetailPage.sizeFields(), CommonUtils.chooseRandomOptionFromSequence(sizeOptions))
    //   .click(driver, ProductdetailPage.buyNowBtn())
    //   .waitForScreenToUpdate(1000000)
    //   .verifyPageTitle(driver, "Store");
  }
  //

}
