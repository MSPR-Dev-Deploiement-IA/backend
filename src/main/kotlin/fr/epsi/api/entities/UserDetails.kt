package fr.epsi.api.entities

import org.springframework.security.core.GrantedAuthority
import org.springframework.security.core.authority.SimpleGrantedAuthority
import org.springframework.security.core.userdetails.UserDetails

class CustomUserDetails(
    private val user: String
) : UserDetails {

    override fun getAuthorities(): Collection<GrantedAuthority> {
        return user.roles.map { SimpleGrantedAuthority(it.name) }
    }

    override fun isEnabled() = true
    override fun isCredentialsNonExpired() = true
    override fun getPassword() = user.password
    override fun getUsername() = user.username
    override fun isAccountNonExpired() = true
    override fun isAccountNonLocked() = true

    // You can add any other user-specific properties as needed
    fun getUserId() = user.id
}
