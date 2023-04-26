package fr.epsi.api.security

import jakarta.servlet.FilterChain
import jakarta.servlet.http.HttpServletRequest
import jakarta.servlet.http.HttpServletResponse
import org.springframework.security.core.Authentication
import org.springframework.security.core.context.SecurityContextHolder
import org.springframework.web.filter.OncePerRequestFilter
import fr.epsi.api.entities.CustomUserDetails
import fr.epsi.api.repositories.UserRepository
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken

class JwtAuthenticationFilter(
    private val jwtTokenProvider: JwtTokenProvider,
    private val userRepository: UserRepository
) : OncePerRequestFilter() {

    override fun doFilterInternal(
        request: HttpServletRequest,
        response: HttpServletResponse,
        filterChain: FilterChain
    ) {
        val token = getTokenFromRequest(request)

        if (token != null && jwtTokenProvider.validateToken(token)) {
            val authentication: Authentication = tokenToAuthentication(token)
            SecurityContextHolder.getContext().authentication = authentication
        }

        filterChain.doFilter(request, response)
    }

    private fun getTokenFromRequest(request: HttpServletRequest): String? {
        val bearerToken = request.getHeader("Authorization")
        return if (bearerToken != null && bearerToken.startsWith("Bearer ")) {
            bearerToken.substring(7)
        } else null
    }

    private fun tokenToAuthentication(token: String): Authentication {
        val user = userRepository.findByUsername(jwtTokenProvider.getUsername(token))
        val userDetails = CustomUserDetails(user!!)
        return UsernamePasswordAuthenticationToken(userDetails, "", userDetails.authorities)
    }
}
