package fr.epsi.api.services

import fr.epsi.api.entities.User
import fr.epsi.api.repositories.UserRepository
import org.springframework.security.crypto.bcrypt.BCrypt
import org.springframework.stereotype.Service

@Service
class UserService(private val userRepository: UserRepository) {
    fun registerUser(username: String, password: String): User {
        val salt = BCrypt.gensalt()
        val hashedPassword = BCrypt.hashpw(password, salt)
        val user = User(username = username, password = hashedPassword)
        return userRepository.save(user)
    }

    fun getUserById(id: Long): User? {
        return userRepository.findById(id).orElse(null)
    }

    fun getUserByUsername(username: String): User? {
        return userRepository.findByUsername(username)
    }

    fun deleteUserById(id: Long) {
        userRepository.deleteById(id)
    }
}
