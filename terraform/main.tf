resource "azurerm_container_registry_task" "tasks" {
  name                  = "backend-task"
  container_registry_id = var.acr_id

  platform {
    os = "Linux"
  }

  docker_step {
    dockerfile_path      = "Dockerfile"
    context_path         = "https://github.com/MSPR-Dev-Deploiement-IA/backend.git#main"
    context_access_token = var.github_access_token
    image_names          = ["backend"]
  }
}

resource "azurerm_container_registry_task_schedule_run_now" "run" {
  container_registry_task_id = azurerm_container_registry_task.tasks[count.index].id
}
