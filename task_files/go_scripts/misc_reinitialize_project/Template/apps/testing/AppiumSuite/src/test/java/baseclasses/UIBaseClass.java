package baseclasses;

import java.awt.AWTException;
import java.awt.Robot;
import java.awt.event.KeyEvent;


import java.io.FileInputStream;
import java.io.IOException;
import java.net.MalformedURLException;
import java.net.URISyntaxException;
import java.util.Random;
import java.util.concurrent.TimeUnit;

import org.openqa.selenium.By;
import org.openqa.selenium.JavascriptExecutor;
import org.openqa.selenium.Keys;
import org.openqa.selenium.Point;
import org.openqa.selenium.WebDriver;
import org.openqa.selenium.remote.service.DriverService;
import org.openqa.selenium.support.ui.WebDriverWait;
import org.testng.ITestContext;
import org.testng.Reporter;
import org.testng.annotations.AfterMethod;
import org.testng.annotations.AfterSuite;
import org.testng.annotations.AfterTest;
import org.testng.annotations.BeforeMethod;
import org.testng.annotations.BeforeSuite;
import org.testng.annotations.Parameters;


import envs.DevEnv;
import interfaces.PlatformsInterface;
import interfaces.EnvironmentsInterface;
import util.CommonUtils;
import util.E2EUtils;

public class UIBaseClass {

  public  static WebDriver driver;
  public static  DriverService service;
  public static JavascriptExecutor js;
  WebDriverWait wait;
  public static DevEnv env;

  public E2EUtils e2eUtil = E2EUtils.getCommonUtils();


  @Parameters({ "paramEnv", "paramPlatform" })
  @BeforeMethod
  public void startUpE2E(String paramEnv, String paramPlatform) throws IOException {
    switch (paramEnv) {
      case "DEV":
        env = EnvironmentsInterface.DEV;
        break;
      case "DEVLOCAL":
        env = EnvironmentsInterface.DEVLOCAL;
        break;
      case "PREVIEW":
        env = EnvironmentsInterface.PREVIEW;
        break;
      case "PROD":
        env = EnvironmentsInterface.PROD;
        break;
      default:
      //Do nothing
    }



    // initFirebaseConnection();
  }


  @Parameters({ "paramEnv", "paramPlatform" })
  @BeforeMethod()
  public void startWithCleanPlatform(String paramEnv, String paramPlatform,ITestContext context) throws InterruptedException, AWTException,URISyntaxException, MalformedURLException {

    switch (paramPlatform) {
      case "ANDROID":
        service = PlatformsInterface.ANDROID.getService();
        driver = PlatformsInterface.ANDROID.getDriver(false,env);
        break;
      case "IOS":
      service = PlatformsInterface.IOS.getService();
        driver = PlatformsInterface.IOS.getDriver(false,env);
        break;
      default:
      //Do nothing
    }

    js = (JavascriptExecutor) driver;
    Reporter.log(context.getCurrentXmlTest().getName(), true);


  }




  @Parameters({ "paramEnv", "paramPlatform" })
  @AfterMethod(alwaysRun = true)
  public void closeBrowserAfterTest(String paramEnv, String paramPlatform) {
      driver.quit();
      service.stop();
  }

  private void zoomOut(int amnt) throws AWTException {
    Robot robot = new Robot();
    for(int i=0; i<amnt; i++){
      robot.keyPress(KeyEvent.VK_CONTROL);
      robot.keyPress(KeyEvent.VK_MINUS);
      robot.keyRelease(KeyEvent.VK_CONTROL);
      robot.keyRelease(KeyEvent.VK_MINUS);
		}
  }


  private void zoomIn(int amnt) throws AWTException {
    Robot robot = new Robot();
    for(int i=0; i<amnt; i++){
			robot.keyPress(KeyEvent.VK_CONTROL);
      robot.keyPress(KeyEvent.VK_EQUALS);
      robot.keyRelease(KeyEvent.VK_CONTROL);
      robot.keyRelease(KeyEvent.VK_EQUALS);
		}
  }



}
