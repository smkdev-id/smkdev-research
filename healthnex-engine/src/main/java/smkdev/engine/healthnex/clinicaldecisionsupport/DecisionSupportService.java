package smkdev.engine.healthnex.clinicaldecisionsupport;

import org.springframework.stereotype.Service;

@Service
public class DecisionSupportService {

    public String getRecommendations(Long patientId) {
        // Mock recommendation logic
        return "Recommendation for patient " + patientId;
    }
}