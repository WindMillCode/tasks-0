package interfaces;

import platforms.E2EAndroid;
import platforms.E2EIOS;
import platforms.E2EPlatform;


public interface PlatformsInterface {
    E2EPlatform ANDROID = new E2EAndroid();
    E2EPlatform IOS = new E2EIOS();
  }
