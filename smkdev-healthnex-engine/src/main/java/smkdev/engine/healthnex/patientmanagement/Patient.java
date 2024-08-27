package smkdev.engine.healthnex.patientmanagement;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class Patient {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    private String name;
    private String dob;
    private String gender;
    private String address;
    private String contactNumber;

    // getters and setters
}
