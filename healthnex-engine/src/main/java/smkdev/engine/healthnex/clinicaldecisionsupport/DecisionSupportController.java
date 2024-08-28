package smkdev.engine.healthnex.clinicaldecisionsupport;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DecisionSupportController {

    @Autowired
    private DecisionSupportService decisionSupportService;

    @GetMapping("/recommendations/{patientId}")
    public String getRecommendations(@PathVariable Long patientId) {
        return decisionSupportService.getRecommendations(patientId);
    }
}