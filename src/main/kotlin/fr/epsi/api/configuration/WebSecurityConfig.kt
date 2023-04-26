package fr.epsi.api.configuration


import fr.epsi.api.repositories.UserRepository
import fr.epsi.api.security.JwtAuthenticationFilter
import fr.epsi.api.security.JwtTokenProvider
import fr.epsi.api.services.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.config.http.SessionCreationPolicy
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.security.web.SecurityFilterChain


@Configuration
class SecurityConfiguration {
    @Autowired
    private lateinit var jwtTokenProvider: JwtTokenProvider

    @Autowired
    private lateinit var userDetailsService: UserService

    @Autowired
    private lateinit var passwordEncoder: BCryptPasswordEncoder

    @Bean
    @Throws(Exception::class)
    fun filterChain(http: HttpSecurity): SecurityFilterChain {
        return http
            .csrf().disable()
            .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS)
            .and().build()
    }

    private fun jwtAuthenticationFilter(): JwtAuthenticationFilter {
        return JwtAuthenticationFilter(jwtTokenProvider, UserRepository())
    }

}

