import { AppBar, Avatar, Box, Button, Chip, Toolbar } from "@mui/material";
import Image from "next/image";
import Link, { LinkProps } from "next/link";
import { useRouter } from "next/router";
import { PropsWithChildren } from "react";
import { useHttp } from "../hooks/useHttp";
import { fetcherStats, httpStats } from "../util/http";

export type NavbarItemProps = LinkProps & { showUnderline: boolean };

export const NavbarItem = (props: PropsWithChildren<NavbarItemProps>) => {
  const { showUnderline, ...linkProps } = props;
  return (
    //@ts-expect-error
    <Button
      component={Link}
      sx={{
        color: "white",
        display: "inline-block",
        textAlign: "center",
        "&::after": (theme) => ({
          content: '""',
          borderBottom: showUnderline
            ? `4px solid ${theme.palette.primary.main}`
            : `4px solid transparent`,
          width: "100%",
          display: "block",
        }),
      }}
      {...linkProps}
    />
  );
};

export const Navbar = () => {
  const router = useRouter();

  const { data, error } = useHttp(
    "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/balance",
    fetcherStats, 
    {
      refreshInterval: 5000
    }
  );

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static" sx={{ background: "none", boxShadow: "none" }}>
        <Toolbar>
          <Image
            src="/img/logo.png"
            width={315}
            height={58}
            alt="logo"
            priority={true}
          />
          <Box sx={{ flexGrow: 1, ml: (theme) => theme.spacing(4) }}>
            <NavbarItem href="/" showUnderline={router.pathname === "/"}>
              Home
            </NavbarItem>
            <NavbarItem
              href="/players"
              showUnderline={router.pathname === "/players"}
            >
              Escalação
            </NavbarItem>
            <NavbarItem
              href="/matches"
              showUnderline={["/matches", "/matches/[id]"].includes(
                router.pathname
              )}
            >
              Jogo
            </NavbarItem>
          </Box>
          <Chip
            label={data ? data.balance : 0}
            avatar={<Avatar>C$</Avatar>}
            color="secondary"
          />
        </Toolbar>
      </AppBar>
    </Box>
  );
};
