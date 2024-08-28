package smkdev.engine.healthnex.medicalrecords;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/medicalrecords")
public class MedicalRecordController {
    @Autowired
    private MedicalRecordRepository medicalRecordRepository;

    @GetMapping
    public List<MedicalRecord> getAllRecords() {
        return medicalRecordRepository.findAll();
    }

    @PostMapping
    public MedicalRecord createRecord(@RequestBody MedicalRecord record) {
        return medicalRecordRepository.save(record);
    }
}