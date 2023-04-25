package fr.epsi.api.security

import io.jsonwebtoken.Jwts
import io.jsonwebtoken.SignatureAlgorithm
import io.jsonwebtoken.security.Keys
import org.springframework.security.core.Authentication
import org.springframework.security.core.userdetails.UserDetails
import org.springframework.stereotype.Component
import java.security.Key
import java.util.*

@Component
class JwtTokenProvider {

    private val key: Key = Keys.secretKeyFor(SignatureAlgorithm.HS256)
    private val validityInMilliseconds: Long = 3600000 // 1 hour

    fun createToken(authentication: Authentication): String {
        val userDetails = authentication.principal as UserDetails
        val now = Date()
        val expiration = Date(now.time + validityInMilliseconds)

        return Jwts.builder()
            .setSubject(userDetails.username)
            .setIssuedAt(now)
            .setExpiration(expiration)
            .signWith(key)
            .compact()
    }

    fun getUsername(token: String): String {
        return Jwts.parserBuilder().setSigningKey(key).build().parseClaimsJws(token).body.subject
    }

    fun validateToken(token: String): Boolean {
        try {
            Jwts.parserBuilder().setSigningKey(key).build().parseClaimsJws(token)
            return true
        } catch (e: Exception) {
            // Log the exception if needed
        }
        return false
    }
}
