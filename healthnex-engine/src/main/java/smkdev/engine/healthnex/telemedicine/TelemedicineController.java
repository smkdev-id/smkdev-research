package smkdev.engine.healthnex.telemedicine;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/telemedicine")
public class TelemedicineController {
    @Autowired
    private TelemedicineRepository telemedicineRepository;

    @GetMapping
    public List<TelemedicineSession> getAllSessions() {
        return telemedicineRepository.findAll();
    }

    @PostMapping
    public TelemedicineSession createSession(@RequestBody TelemedicineSession session) {
        return telemedicineRepository.save(session);
    }
}
