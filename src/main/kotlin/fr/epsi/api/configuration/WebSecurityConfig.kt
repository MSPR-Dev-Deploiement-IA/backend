package fr.epsi.api.configuration

import fr.epsi.api.security.JwtAuthenticationFilter
import fr.epsi.api.security.JwtTokenProvider
import fr.epsi.api.services.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter
import org.springframework.security.config.http.SessionCreationPolicy
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter

@Configuration
@EnableWebSecurity
class WebSecurityConfig(
    private val userService: UserService,
    private val jwtTokenProvider: JwtTokenProvider
) : WebSecurityConfigurerAdapter() {

    @Autowired
    fun configureGlobal(auth: AuthenticationManagerBuilder) {
        auth.userDetailsService(userService).passwordEncoder(BCryptPasswordEncoder())
    }

    override fun configure(http: HttpSecurity) {
        http
            .csrf().disable()
            .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS)
            .and()
            .authorizeRequests()
            .antMatchers("/auth/*").permitAll()
            .anyRequest().authenticated()
            .and()
            .addFilterBefore(
                JwtAuthenticationFilter(jwtTokenProvider),
                UsernamePasswordAuthenticationFilter::class.java
            )
    }
}
