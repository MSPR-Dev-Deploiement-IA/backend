package fr.epsi.api.controllers

import org.springframework.http.ResponseEntity
import org.springframework.security.access.prepost.PreAuthorize
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/secure")
class SecureController {

    @GetMapping
    @PreAuthorize("hasRole('ROLE_USER')")
    fun getSecureData(): ResponseEntity<String> {
        val data = "This data is secure"
        return ResponseEntity.ok(data)
    }
}