import { Box, styled, Typography } from "@mui/material";
import { match } from "assert";
import Image from "next/image";
import { Match, TeamsImagesMap } from "../util/models";

const ResultContainer = styled(Box)(({ theme }) => ({
  display: "flex",
  width: "400px",
  backgroundColor: theme.palette.background.default,
  alignItems: "center",
  padding: 0,
  border: 'none !important',
  boxShadow: 'none'
}));

const ResultItem = styled(Box)(({ theme }) => ({
  height: "55px",
  display: "flex",
  alignItems: "center",
}));

type FlagProps = {
  src: string;
  alt: string;
};

const Flag = (props: FlagProps) => {
  return <Image src={props.src} alt={props.alt} width={121} height={76} style={{
    marginLeft: '-5px',
    marginRight: '-5px'
  }} />;
};

type MatchResultProps = {
  match: Match;
};

export const MatchResult = (props: MatchResultProps) => {
  const { match } = props;
  return (
    <Box display="flex">
      <Flag src={TeamsImagesMap[match.team_a]} alt={match.team_a} />
      <ResultContainer>
        <ResultItem width={"150px"} justifyContent="flex-end">
          <Typography variant="h6">Brasil</Typography>
        </ResultItem>
        <ResultItem width={"100px"} justifyContent="center" position="relative">
          <Box sx={{ position: "absolute", top: 0, fontSize: "0.70rem" }}>
            12/12/2022 00:00
          </Box>
          <Typography variant="h6" sx={{ fontWeight: "900" }}>
            1-0
          </Typography>
        </ResultItem>
        <ResultItem width={"150px"} justifyContent="flex-start">
          <Typography variant="h6">Argentina</Typography>
        </ResultItem>
      </ResultContainer>
      <Flag src={TeamsImagesMap[match.team_b]} alt={match.team_a} />
    </Box>
  );
};
