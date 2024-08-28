package smkdev.engine.healthnex.analyticsreporting;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
public class AnalyticsController {

    @Autowired
    private AnalyticsService analyticsService;

    @GetMapping("/dashboard")
    public Map<String, Object> getDashboardData() {
        return analyticsService.getDashboardData();
    }
}