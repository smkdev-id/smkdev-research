package smkdev.engine.healthnex.telemedicine;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import java.time.LocalDateTime;

@Entity
public class TelemedicineSession {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    private Long patientId;
    private Long doctorId;
    private LocalDateTime sessionDate;
    private String sessionLink;

    // getters and setters
}
