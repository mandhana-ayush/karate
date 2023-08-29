package hooks;

import com.intuit.karate.RuntimeHook;
import com.intuit.karate.core.FeatureRuntime;

import helpers.DbHandler;

public class KarateHooks implements com.intuit.karate.RuntimeHook{
  public void afterFeature(FeatureRuntime fr){
    System.out.println("Feature is finished: " );
    
    DbHandler.cleanup();
    RuntimeHook.super.afterFeature(fr);
  }
}
