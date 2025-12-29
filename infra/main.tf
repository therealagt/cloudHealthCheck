provider "google" {
  project = "go-healthcheck-test-291225"
  region  = "us-central1"
  zone    = "us-central1-a"
}

# 1. Die "Guten" VMs (Compliant)
# Diese haben Labels, wie es sich gehört.
resource "google_compute_instance" "compliant_vms" {
  count        = 3
  name         = "compliant-vm-${count.index}"
  machine_type = "e2-micro" # Günstigste Option

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }

  # Hier sind die Labels definiert
  labels = {
    environment = "production"
    team        = "backend"
    managed_by  = "terraform"
  }
}

# 2. Die "Schlechten" VMs (Non-Compliant)
# Diese haben KEINE Labels -> Das soll dein Tool finden!
resource "google_compute_instance" "rogue_vms" {
  count        = 2
  name         = "rogue-vm-${count.index}"
  machine_type = "e2-micro"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }

  # FEHLER: Keine Labels definiert!
}