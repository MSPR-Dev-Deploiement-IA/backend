package fr.epsi.api.services

import fr.epsi.api.entities.Role
import fr.epsi.api.entities.User
import fr.epsi.api.repositories.UserRepository
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.stereotype.Service

@Service
class UserService(private val userRepository: UserRepository, private val passwordEncoder: BCryptPasswordEncoder) {

    fun registerUser(username: String, password: String, roles: Set<Role>): User {
        val hashedPassword = passwordEncoder.encode(password)
        val newUser = User(username = username, password = hashedPassword, roles = roles)
        return userRepository.save(newUser)
    }
}
