export type SiteConfig = typeof siteConfig;

export const siteConfig = {
  name: "OpenWeightlifting",
  description: "OpenWeightlifting is an open source weightlifting knowledge platform.",
  navItems: [
    {
      label: "Home",
      href: "/",
    },
    {
      label: "Sinclair Calculator",
      href: "/sinclair",
    }
  ],
  navMenuItems: [
    {
      label: "Home",
      href: "/",
    },
    {
      label: "Sinclair Calculator",
      href: "/sinclair",
    },
  ],
  links: {
    github: "https://github.com/euanwm/OpenWeightlifting",
    instagram: "https://www.instagram.com/openweightlifting/",
    discord: "https://discord.com/kqnBqdktgr"
  },
};