package platforms;

import java.net.MalformedURLException;
import java.net.URI;
import java.net.URISyntaxException;

import org.openqa.selenium.WebDriver;
import org.openqa.selenium.remote.service.DriverService;

import envs.DevEnv;
import io.appium.java_client.android.AndroidDriver;
import io.appium.java_client.android.options.UiAutomator2Options;
import io.appium.java_client.service.local.AppiumDriverLocalService;

public class E2EAndroid extends E2EPlatform {



  public WebDriver getDriver(Boolean headless,DevEnv env) throws InterruptedException, URISyntaxException, MalformedURLException {

    UiAutomator2Options options = new UiAutomator2Options()
        .setUdid("emulator-5554")
        .setAppPackage(env.appPacakge)
        .setAppActivity(env.appActivity);
    driver = new AndroidDriver(
        new URI(env.appiumServerURL).toURL(), options);
    return driver;
  }

  public DriverService getService() {
    AppiumDriverLocalService service = AppiumDriverLocalService.buildDefaultService();
    service.start();
    return service;
  }
}
