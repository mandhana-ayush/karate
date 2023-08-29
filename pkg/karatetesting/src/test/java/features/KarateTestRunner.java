// package features;

// import com.intuit.karate.junit5.Karate;

// import hooks.KarateHooks;

// import org.junit.runner.RunWith;

// public class KarateTestRunner {
//   @Karate.Test

//   public Karate testFeature1(){
//     System.out.println("testfeature1");
//     // return Karate.run("classpath:features").hook(new KarateHooks()).relativeTo(getClass());
//     return Karate.run("classpath:features").relativeTo(getClass());
//   }

  // @Karate.Test

  // public Karate testFeature2(){
  //   return Karate.run("classpath:features/opd").hook(new KarateHooks()).relativeTo(getClass())
  // }

  // @Karate.Test
  // public Karate testFeature2(){
  //   System.out.println("testFeature2");
  //   return Karate.run("my").hook(new KarateHooks()).relativeTo(getClass());
  // }
  // @Test
// }


package features;

import com.intuit.karate.Results;
import com.intuit.karate.Runner;
import com.intuit.karate.junit5.Karate;

import hooks.KarateHooks;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class KarateTestRunner {
    
    @Test
    public void testFeature1() {  
        String[] tags = {"@smoke"};

        System.out.println("testFeature1");
        Results result = Runner.path("classpath:features").tags(System.getProperty("karateOptions")).parallel(Integer.parseInt(System.getProperty("threadCount", "1")));
        assertEquals(0, result.getFailCount(),"Failures found in test run");
    }
}


