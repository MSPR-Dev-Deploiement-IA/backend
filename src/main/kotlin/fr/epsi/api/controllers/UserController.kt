package fr.epsi.api.controllers

import fr.epsi.api.entities.Role
import fr.epsi.api.entities.User
import fr.epsi.api.repositories.UserRepository
import fr.epsi.api.services.UserService
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.security.crypto.bcrypt.BCrypt
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/auth")
class UserController(
    private val userRepository: UserRepository,
    private val userService: UserService,
) {

    data class Credentials(
        val username: String,
        val password: String,
        val roles: Set<Role>
    )

    @PostMapping("/login")
    fun login(@RequestBody credentials: Credentials): ResponseEntity<Any> {
        val user = userRepository.findByUsername(credentials.username)
            ?: return ResponseEntity.status(HttpStatus.UNAUTHORIZED).build()

        if (BCrypt.checkpw(credentials.password, user.password)) {
            return ResponseEntity.ok(user)
        } else {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).build()
        }
    }

    @PostMapping("/register")
    fun register(@RequestBody credentials: Credentials): ResponseEntity<User> {
        val user = userService.registerUser(credentials.username, credentials.password, credentials.roles)
        return ResponseEntity.status(HttpStatus.CREATED).body(user)
    }
}

