import { Box } from "@mui/material";
import type { NextPage } from "next";
import { useRouter } from "next/router";
import { MatchResult } from "../../components/MatchResult";
import { Page } from "../../components/Page";
import { useHttp } from "../../hooks/useHttp";
import { fetcherStats } from "../../util/http";
import { Match } from "../../util/models";

const ListMatchesPage: NextPage = () => {
  const { data } = useHttp<Match[]>("/matches", fetcherStats, {
    refreshInterval: 5000,
  });
  const router = useRouter();
  return (
    <Page>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          flexDirection: "column",
          gap: (theme) => theme.spacing(3),
        }}
      >
        {data &&
          data.map((match, key) => (
            <Box
              key={key}
              sx={{ cursor: "pointer" }}
              onClick={() => router.push(`/matches/${match.id}`)}
            >
              <MatchResult match={match} />
            </Box>
          ))}
      </Box>
    </Page>
  );
};

export default ListMatchesPage;