package hooks;

import com.intuit.karate.Results;
import com.intuit.karate.Runner;
import com.intuit.karate.RuntimeHook;
import com.intuit.karate.core.ScenarioRuntime;
import com.intuit.karate.junit5.Karate;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class KarateHooks implements com.intuit.karate.RuntimeHook{

    @Override
    public boolean beforeScenario(ScenarioRuntime sc) {
        // This method will be executed before each scenario
        System.out.println("Before scenario hook executed");

        RuntimeHook.super.beforeScenario(sc);
        return true;
    }
}
