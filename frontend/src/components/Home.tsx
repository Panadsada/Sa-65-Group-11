// import { createStyles, makeStyles, Theme } from "@mui/material/styles";
import Container from "@mui/material/Container";

// const useStyles = makeStyles((theme: Theme) =>
//   createStyles({
//     container: {
//       marginTop: theme.spacing(2),
//     },
//     table: {
//       minWidth: 650,
//     },
//     tableSpace: {
//       marginTop: 20,
//     },
//   })
// );

function Home() {
  // const classes = useStyles();

  return (
    <div>
      <Container sx={{ marginTop: 2}} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบบันทึกอาการโรงพยาบาลมหาวิทยาลัยเทคโนโลยีสุรนารี</h1>
        <h3>Requirements</h3>
        <p>
          ระบบบันทึกอาการเป็นระบบที่ทำมาเพื่อให้แพทย์ที่ทำการรักษาที่โรงพยาบาลจะทราบถึงข้อมูล
          ของผู้ป่วยจะทำการบันทึกเกี่ยวกับอาการเจ็บป่วย พบว่าแพทย์ต้องทำการเขียนบันทึกลงกระดาษ
          อาจจะทำให้มีเกิดเอกสารตกหล่นหรือสูญหายได้ แพทย์จึงมีความต้องการที่จะให้มีระบบบันทึกอาการ
          ของผู้ป่วยเพื่อที่จะมีระบบแก้ปัญหาดังที่กล่าวไว้ข้างต้น เราจึงได้ออกแบบระบบบันทึกอาการของผู้ป่วย
          เป็นระบบที่ให้แพทย์ใช้ ระบบซึ่งเป็นแพทย์ที่ทำการรักษาในโรงพยาบาลสามารถ login เข้าระบบเพื่อ
          กรอกข้อมูลอาการของผู้ป่วย โดยแพทย์ผู้ใช้งานแต่ละคนสามารถบันทึกอาการของคนไข้แต่ละคน 
          เลือกบริเวณที่เจ็บป่วยได้ เลือกแผนกที่ตรวจ เลือกวันที่เวลาที่ตรวจ และสามารถที่จะกรอกอาการ
          เพิ่มเติมได้ จึงทำให้สะดวกสบายต่อการบันทึกมากขึ้น

        </p>
      </Container>
    </div>
  );
}
export default Home;