package smkdev.engine.healthnex.analyticsreporting;

import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
public class AnalyticsService {

    public Map<String, Object> getDashboardData() {
        // Mock dashboard data
        Map<String, Object> data = new HashMap<>();
        data.put("totalPatients", 100);
        data.put("appointmentsToday", 10);
        return data;
    }
}