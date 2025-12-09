use sycamore::prelude::*;

#[derive(Clone)]
pub enum VaultRoutes {
    Home,
    NotFound,
}

pub fn VaultPage() -> View {
    view! {
        h1 { "Vault Page" }
    }
}