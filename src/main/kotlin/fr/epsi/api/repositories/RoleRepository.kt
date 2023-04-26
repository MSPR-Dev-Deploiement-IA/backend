package fr.epsi.api.repositories

import fr.epsi.api.entities.Role
import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository

@Repository
interface RoleRepository : CrudRepository<Role, Long> {
    fun findByName(name: String): Role?
}
