use sycamore::prelude::*;
use crate::components::navbar::Navbar;

#[component]
pub fn IndexPage() -> View {
    view! {
        // Page wrapper: full viewport height, column layout
        div(class="min-h-screen flex flex-col bg-base-100") {
            // Navbar (keeps its own markup)
            Navbar()

            // Main content area that will get the scrollbar when content overflows.
            main(class="flex-1 overflow-y-auto") {
                // Hero section
                div(class="hero py-12") {
                    div(class="hero-content w-full max-w-7xl mx-auto") {
                        div(class="grid grid-cols-1 md:grid-cols-2 gap-8 items-center") {
                            // Left: Title, subtitle, CTA buttons and badges
                            div(class="text-center md:text-left space-y-6") {
                                h1(class="text-5xl font-bold") { "PepeVault" }
                                p(class="max-w-md text-lg text-muted") {
                                    "Password Manager made in Rust"
                                }

                                div(class="flex flex-col sm:flex-row gap-3 justify-center md:justify-start") {
                                    a(class="btn btn-primary btn-lg", href="/login") { "Get started" }
                                    a(class="btn btn-ghost btn-lg", href="#") { "GitHub" }
                                }
                            }

                            // Right: abstract composition (image)
                            // - container uses bg-base-100 so the icon sits on that color
                            // - `md:justify-end` keeps it to the right on larger screens
                            div(class="flex items-center md:justify-end") {
                                // Card-like wrapper so the icon has padding, rounded corners, and fits well
                                div(class="w-full max-w-lg p-3 rounded-xl shadow-lg flex items-center justify-center") {
                                    img(class="w-full h-auto object-contain", src="public/logo.webp", alt="PepeVault icon", loading="lazy")
                                }
                            }
                        } // end grid
                    } // end hero-content
                } // end hero
            } // end main
        } // end wrapper
    }
}
