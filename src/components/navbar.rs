use sycamore::prelude::*;

#[component]
pub fn Navbar() -> View {
    view! {
        div (class="navbar justify-between bg-base-300") {
            // Logo
            div (class="navbar-start") {
                div (class="avatar") {
                    div (class="w-15 rounded-full") {
                        img (src="public/cat.webp") { }
                    }
                }
            }

            // Mobile
            div (class="navbar-end dropdown dropdown-end sm:hidden") {
                label (tabindex="0", class="btn btn-ghost") {
                    i (class="fa-solid fa-bars text-lg") { }
                }

                ul (tabindex="0", class="dropdown-content menu z-[1] bg-base-200 p-6 rounded-box shadow w-56 gap-2") {
                    li {
                        a (class="btn btn-sm btn-primary") { "Signup" }
                    }
                    li {
                        a (class="btn btn-sm btn-primary") { "Login" }
                    }
                }
            }

            // Desktop
            div (class="navbar-end") {
                ul (class="hidden sm:flex sm:menu-horizontal gap-2") {
                    li {
                        a (class="btn btn-primary", href="/signup") { "Signup" }
                    }
                    li {
                        a (class="btn btn-primary", href="/login") { "Login" }
                    }
                }
            }
        }
    }
}