import GroupsIcon from "@mui/icons-material/Groups";
import { Button, Divider, Grid, styled } from "@mui/material";
import type { NextPage } from "next";
import Link from "next/link";
import { Label } from "../components/Label";
import { Page } from "../components/Page";
import { Section } from "../components/Section";
import { TeamLogo } from "../components/TeamLogo";
import { useHttp } from "../hooks/useHttp";
import { fetcherStats } from "../util/http";

const BudgetContainer = styled(Section)(({ theme }) => ({
  width: "800px",
  height: "300px",
  marginTop: theme.spacing(8),
  display: "flex",
  alignItems: "center",
}));

const HomePage: NextPage = () => {
  const { data, error } = useHttp(
    "/my-teams/22087246-01bc-46ad-a9d9-a99a6d734167/balance",
    fetcherStats,
    {
      refreshInterval: 5000,
    }
  );

  return (
    <Page>
      <Grid
        container
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: (theme) => theme.spacing(3),
        }}
      >
        <Grid item>
          <TeamLogo
            sx={{ position: "absolute", left: 0, right: 0, m: "auto" }}
          />
          <BudgetContainer>
            <Grid container>
              <Grid
                item
                xs={5}
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Label>Última pontuação</Label>
                <Label>-</Label>
              </Grid>
              <Grid
                item
                xs={2}
                sx={{ display: "flex", justifyContent: "center" }}
              >
                <Divider orientation="vertical" sx={{ height: "auto" }} />
              </Grid>
              <Grid
                item
                xs={5}
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Label>Patrimônio</Label>
                <Label>{data ? data.balance : 0}</Label>
              </Grid>
            </Grid>
          </BudgetContainer>
        </Grid>
        <Grid item>
          <Button
            component={Link}
            href="/players"
            variant="contained"
            startIcon={<GroupsIcon />}
          >
            Escalar Jogadores
          </Button>
        </Grid>
      </Grid>
    </Page>
  );
};

export default HomePage;
