package app;

import org.testng.annotations.DataProvider;

import interfaces.*;

public class CommonDataProviders {

  @DataProvider
  public static Object[][] environments() {
    return new Object[][] {
        { EnvironmentsInterface.DEV },
        { EnvironmentsInterface.PREVIEW },
        { EnvironmentsInterface.PROD }
    };
  }

  @DataProvider
  public static Object[][] browsers() {
    return new Object[][] {
        { PlatformsInterface.ANDROID },
        { PlatformsInterface.IOS }
    };
  }
}
