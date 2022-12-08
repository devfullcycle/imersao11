import { Box, Button } from "@mui/material";
import type { NextPage } from "next";
import { MatchResult } from "../../components/MatchResult";
import { Page } from "../../components/Page";

const ListMatchesPage: NextPage = () => {
  return (
    <Page>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          flexDirection: 'column',
          gap: (theme) => theme.spacing(3),
        }}
      >
        <MatchResult match={{ team_a: "Brasil", team_b: "Argentina" }} />
        <MatchResult match={{ team_a: "Brasil", team_b: "Argentina" }} />
        <MatchResult match={{ team_a: "Brasil", team_b: "Argentina" }} />
      </Box>
    </Page>
  );
};

export default ListMatchesPage;
