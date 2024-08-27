package smkdev.engine.healthnex.patientportal;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PatientPortalController {

    @GetMapping("/portal")
    public String portalHome() {
        return "Welcome to the Patient Portal!";
    }
}