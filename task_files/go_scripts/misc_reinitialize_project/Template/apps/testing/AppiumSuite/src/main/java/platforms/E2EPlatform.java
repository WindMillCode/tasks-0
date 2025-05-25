package platforms;

import java.net.MalformedURLException;
import java.net.URISyntaxException;

import org.openqa.selenium.WebDriver;
import org.openqa.selenium.remote.service.DriverService;

import envs.DevEnv;

public class E2EPlatform {
  public static WebDriver driver;
  public static DriverService service;

  public  WebDriver getDriver(Boolean headless,DevEnv env) throws InterruptedException, URISyntaxException, MalformedURLException {
    return driver;

  }

  public DriverService getService() {
    return service;
  }

}
