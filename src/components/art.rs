use sycamore::prelude::*;

#[component]
pub fn RustLockArt() -> View {
    view! {
        div(class="flex items-center justify-center", aria-hidden="true") {
            div(class="relative w-80 h-80 rounded-full bg-gradient-to-br from-orange-400 via-amber-500 to-rose-400 shadow-2xl flex items-center justify-center") {
                div(class="absolute inset-0 rounded-full ring-1 ring-inset ring-white/10 pointer-events-none") {}

                // BIGGER GEAR
                svg(class="w-72 h-72", viewBox="0 0 100 100", xmlns="http://www.w3.org/2000/svg") {
                    circle(cx="50", cy="50", r="36", fill="rgba(255,255,255,0.04)") {}

                    polygon(points="80,50 65,75 35,75 20,50 35,25 65,25",
                            fill="none",
                            stroke="rgba(255,255,255,0.92)",
                            stroke-width="3.2",
                            stroke-linejoin="round") {}

                    circle(cx="50", cy="50", r="10.5", fill="rgba(0,0,0,0.28)", stroke="rgba(255,255,255,0.62)", stroke-width="1.4") {}

                    rect(x="76", y="45", width="6.6", height="10", rx="1.2", transform="rotate(18 79.3 50)", fill="rgba(255,255,255,0.92)") {}
                    rect(x="62", y="72", width="6.6", height="10", rx="1.2", transform="rotate(78 65.3 77)", fill="rgba(255,255,255,0.92)") {}
                    rect(x="32", y="72", width="6.6", height="10", rx="1.2", transform="rotate(138 35.3 77)", fill="rgba(255,255,255,0.92)") {}
                    rect(x="13.8", y="45", width="6.6", height="10", rx="1.2", transform="rotate(198 17.1 50)", fill="rgba(255,255,255,0.92)") {}
                    rect(x="32", y="20.5", width="6.6", height="10", rx="1.2", transform="rotate(258 35.3 25.5)", fill="rgba(255,255,255,0.92)") {}
                    rect(x="62", y="20.5", width="6.6", height="10", rx="1.2", transform="rotate(318 65.3 25.5)", fill="rgba(255,255,255,0.92)") {}

                    circle(cx="50", cy="50", r="24", fill="none", stroke="rgba(255,255,255,0.04)", stroke-width="6") {}
                }

                // BIGGER AND RECENTERED LOCK
                div(class="absolute drop-shadow-2xl -translate-y-0.5") {
                    svg(class="w-40 h-40", viewBox="0 0 64 64", xmlns="http://www.w3.org/2000/svg") {
                        defs {
                            linearGradient(id="lockGrad", x1="0%", x2="0%", y1="0%", y2="100%") {
                                stop(offset="0%", stop-color="#2b2b2b", stop-opacity="0.95") {}
                                stop(offset="100%", stop-color="#0c0c0c", stop-opacity="0.9") {}
                            }
                            linearGradient(id="shackleGrad", x1="0%", x2="0%", y1="0%", y2="100%") {
                                stop(offset="0%", stop-color="#ffffff", stop-opacity="0.95") {}
                                stop(offset="100%", stop-color="#ffffff", stop-opacity="0.75") {}
                            }
                        }

                        ellipse(cx="32", cy="50.5", rx="14", ry="3.5", fill="rgba(0,0,0,0.35)") {}

                        rect(x="12.5", y="24.5", width="39", height="26", rx="4.5",
                             fill="url(#lockGrad)", stroke="rgba(255,255,255,0.08)", stroke-width="0.9") {}

                        path(d="M22.5 27.5 C22.5 18.8, 41.5 18.8, 41.5 27.5",
                             fill="none", stroke="rgba(255,255,255,0.95)", stroke-width="3.2",
                             stroke-linecap="round") {}

                        path(d="M23.6 27.5 C23.6 19.3, 40.4 19.3, 40.4 27.5",
                             fill="none", stroke="rgba(255,255,255,0.55)", stroke-width="1.2",
                             stroke-linecap="round") {}

                        circle(cx="32", cy="36.5", r="3.6", fill="rgba(255,255,255,0.96)") {}
                        path(d="M31.2 40.2 L32 42.6 L32.8 40.2 Z", fill="rgba(255,255,255,0.96)") {}

                        path(d="M15 27 C22 24, 36 24, 48 27", stroke="rgba(255,255,255,0.06)",
                             stroke-width="6", stroke-linecap="round") {}

                        circle(cx="32", cy="36.5", r="1.1", fill="rgba(0,0,0,0.18)", transform="translate(0,0.8)") {}
                    }
                }
            }
        }
    }
}
