package smkdev.engine.healthnex.telemedicine;

import org.springframework.data.jpa.repository.JpaRepository;

public interface TelemedicineRepository extends JpaRepository<TelemedicineSession, Long> {
}
