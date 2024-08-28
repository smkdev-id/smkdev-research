package smkdev.engine.healthnex.billinginsurance;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/billing")
public class BillingController {
    @Autowired
    private BillingRepository billingRepository;

    @GetMapping
    public List<Billing> getAllBilling() {
        return billingRepository.findAll();
    }

    @PostMapping
    public Billing createBilling(@RequestBody Billing billing) {
        return billingRepository.save(billing);
    }
}
