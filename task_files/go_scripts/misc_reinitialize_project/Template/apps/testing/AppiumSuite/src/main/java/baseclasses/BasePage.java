package baseclasses;

import org.openqa.selenium.By;


public class BasePage {

  public  static String generateCSSSelector(String prefix,String suffix){
    return prefix+"_"+suffix;
  }

  public static By getByWMLButton(String id){
    return By.cssSelector("#"+id+ " button");
  }
  
}
